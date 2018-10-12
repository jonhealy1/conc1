import java.util.Random;
import java.util.concurrent.BrokenBarrierException;
import java.util.concurrent.CountDownLatch;
import java.util.concurrent.CyclicBarrier;
import java.util.concurrent.Semaphore;

public class SmokingAgent extends Thread {
 
    public Semaphore semaphoreSmoked = new Semaphore(0);
    public Semaphore semaphoreIngredient = new Semaphore(0);
    public String disposedIngredients = new String();
    public CountDownLatch latch;
    public Semaphore semaphoreLatchStart = new Semaphore(0);
    public boolean finished = false;
 
    public void run() {
        Random random = new Random();
        int currentIngredients;
        for (int i = 0; i < 1000; i++) {
            latch = new CountDownLatch(3);
            semaphoreLatchStart.release(3);
            try {
                latch.await();
            } catch (InterruptedException e1) {
                e1.printStackTrace();
            }
            currentIngredients = random.nextInt(3);
            if (currentIngredients == 0) {
                disposedIngredients = "paper and a match";
            }
            if (currentIngredients == 1) {
                disposedIngredients = "tobacco and a match";
            }
            if (currentIngredients == 2) {
                disposedIngredients = "paper and tobacco";
            }
            System.out.println("agent provides " + disposedIngredients);
            semaphoreIngredient.release(3);
            try {
                semaphoreSmoked.acquire(3);
            } catch (InterruptedException e) {
                // TODO Auto-generated catch block
                e.printStackTrace();
            }
        }
        finished = true;
        semaphoreLatchStart.release(3);
        semaphoreIngredient.release(3);
    }
}