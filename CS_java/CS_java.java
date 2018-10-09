//https://paulamarcu.wordpress.com/2014/08/16/cigarette-smokers-problem/

/**Having four threads: an agent and three smokers; the smokers wait for ingredients
as long as the agent is available, then take the missing ingredients,
make and smoke the cigarettes. The ingredients are tobacco, paper and matches.
The agent has an infinite supply of all three ingredients and
each smoker has an infinite supply of one of the ingredients; that is, one smoker
has matches, another has paper and the third has tobacco.
The agent offers the missing ingredients for a given number of times and the smokers try
to smoke as long as the agent has more ingredients to offer.
The agent repeatedly chooses two different ingredients at random and makes
them available to the smokers. Depending on which ingredients are chosen, the
smoker with the complementary ingredient should pick up both resources and
proceed.
For example, if the agent puts out tobacco and paper, the smoker with the
matches should pick up both ingredients, make a cigarette and smoke.
*/

//package ro.paula.exercises.smoke;
 
import java.util.Random;
import java.util.concurrent.BrokenBarrierException;
import java.util.concurrent.CountDownLatch;
import java.util.concurrent.CyclicBarrier;
import java.util.concurrent.Semaphore;
 

 

 
public class CS_java {
 
    public static void main(String[] args) {
        // TODO Auto-generated method stub
        SmokingAgent agent = new SmokingAgent();
        Smoker tabaccoSmoker = new Smoker(agent, "Tabacco");
        Smoker paperSmoker = new Smoker(agent, "Paper");
        Smoker matchesSmoker = new Smoker(agent, "Matches");
        agent.start();
        tabaccoSmoker.start();
        paperSmoker.start();
        matchesSmoker.start();
    }
}