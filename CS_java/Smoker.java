import java.util.Random;
import java.util.concurrent.BrokenBarrierException;
import java.util.concurrent.CountDownLatch;
import java.util.concurrent.CyclicBarrier;
import java.util.concurrent.Semaphore;

public class Smoker extends Thread {
 
    private SmokingAgent agent;
    private String ownIngredient;
    private String missingIngredients;
 
    public Smoker(SmokingAgent agent, String ingredient) {
        this.agent = agent;
        ownIngredient = ingredient;
        if (ownIngredient.equals("Tabacco")) {
            missingIngredients = "Paper and Matches";
        }
        if (ownIngredient.equals("Paper")) {
            missingIngredients = "Tabacco and Matches";
        }
        if (ownIngredient.equals("Matches")) {
            missingIngredients = "Paper and Tabacco";
        }
    }
 
    public void run() {
        while (true) {
            if (agent.finished) {
                return;
            }
            try {
                agent.semaphoreLatchStart.acquire();
            } catch (InterruptedException e1) {
                e1.printStackTrace();
            }
            synchronized (agent.latch) {
                agent.latch.countDown();
            }
            try {
                agent.semaphoreIngredient.acquire();
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
            if (agent.finished) {
                System.out.println("The smoking agent " +
                        "is no longer available.I'll give up smoking.");
                return;
            }
            if (agent.disposedIngredients.equals(missingIngredients)) {
                System.out.println("I owe " + ownIngredient + " and I received "
                        + agent.disposedIngredients + ", so now I can smoke.");
            }
            agent.semaphoreSmoked.release();
        }
    }
}