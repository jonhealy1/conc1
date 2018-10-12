import java.util.Random;
import java.util.concurrent.BrokenBarrierException;
import java.util.concurrent.CountDownLatch;
import java.util.concurrent.CyclicBarrier;
import java.util.concurrent.Semaphore;
import java.util.concurrent.TimeUnit;

public class Smoker extends Thread {
 
    private SmokingAgent agent;
    private String ownIngredient;
    private String missingIngredients;
 
    public Smoker(SmokingAgent agent, String ingredient) {
        this.agent = agent;
        ownIngredient = ingredient;
        if (ownIngredient.equals("tobacco")) {
            missingIngredients = "paper and a match";
        }
        if (ownIngredient.equals("paper")) {
            missingIngredients = "tobacco and a match";
        }
        if (ownIngredient.equals("matches")) {
            missingIngredients = "paper and tobacco";
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
                System.out.println("Agent " +
                        "not available.");
                return;
            }
            if (agent.disposedIngredients.equals(missingIngredients)) {
                System.out.println("smoker with " + ownIngredient + " makes cigarette ");
                System.out.println("smoker with " + ownIngredient + " smokes ");
            }
            agent.semaphoreSmoked.release();
        }
    }
}