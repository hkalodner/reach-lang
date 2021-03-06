import * as RPS        from './build/rps.mjs';
import { runTests, assert } from '@reach-sh/stdlib/tester.mjs';
import * as stdlib from '@reach-sh/stdlib';

const log = (msg, ret = true) => () => { console.log(`...${msg}`); return ret; };
const interactWithAlice =
      ({ params: log(`params`),
         getHand: log(`Alice getHand`, 'SCISSORS'),
         commits: log(`commits`),
         reveals: (handB) => log(`reveals ${handB}`)(),
         outcome: log(`Alice outcome`) });
const interactWithBob =
      ({ accepts: (wagerAmount, escrowAmount) => log(`accepts ${wagerAmount} ${escrowAmount}`)(),
         getHand: log(`Bob getHand`, 'PAPER'),
         shows: log(`shows`),
         outcome: log(`Bob outcome`) });

runTests(async () => {
  console.log(`Running game...`);

  const wagerInWei = stdlib.toWeiBN('1.5', 'ether');
  const escrowInWei = stdlib.toWeiBN('0.15', 'ether');
  const startingBalance = stdlib.toWeiBN('100', 'ether');

  const alice = await stdlib.newTestAccount(startingBalance);
  const bob = await stdlib.newTestAccount(startingBalance);
  const balanceStartAlice = await stdlib.balanceOf(alice);
  const balanceStartBob = await stdlib.balanceOf(bob);

  const ctcAlice = await alice.deploy(RPS);
  const ctcBob = await bob.attach(RPS, ctcAlice.address,
                                  ctcAlice.creation_block);

  const [ outcomeBob, outcomeAlice ] =
        await Promise.all([
          RPS.B(ctcBob, interactWithBob),
          RPS.A(ctcAlice, interactWithAlice,
                wagerInWei, escrowInWei)]);

  const balanceEndAlice = await stdlib.balanceOf(alice);
  const balanceEndBob = await stdlib.balanceOf(bob);

  assert.deepEqual(outcomeAlice, outcomeBob, `outcomes disagree`);

  assert.deepEqual(outcomeAlice, ['Alice wins']);

  const [ balStartWinner, balStartLoser, balEndWinner, balEndLoser ] =
        [ balanceStartAlice, balanceStartBob, balanceEndAlice, balanceEndBob ];

  assert.ok(balEndWinner.gte(balStartWinner), `winner' >= winner`);
  assert.ok(balEndLoser.lt(balStartLoser), `loser' < loser`);
});
