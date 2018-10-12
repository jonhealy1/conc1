import java.util.concurrent.CountDownLatch;
import java.util.concurrent.Semaphore;
import java.time.Instant;
import java.util.concurrent.TimeUnit;
import java.util.concurrent.*;

/* This code was inspired by github.com/soniakeys/LittleBookOfSemaphores 
    and The Little Book of Semaphores */

class RW_java {
    static Semaphore mutex = new Semaphore(1); //readLock
    static Semaphore roomEmpty = new Semaphore(1); //writeLock
    static int readers = 0;
    static int virtualBytes = 0;

    static class Reader implements Runnable {
        @Override
        public void run() {
            try {
                //Acquire Section
                mutex.acquire();
                readers++;
                if (readers == 1) {
                    roomEmpty.acquire();
                }
                mutex.release();

                //Reading section
                System.out.println(Thread.currentThread().getName() + " sees " + virtualBytes + " bytes");
                mutex.acquire();
                readers--;
                
                if(readers == 0) {
                    roomEmpty.release();
                }
                mutex.release();
            } catch (InterruptedException e) {
                System.out.println(e.getMessage());
            }
        }
    }

    static class Writer implements Runnable {
        @Override
        public void run() {
            try {
                roomEmpty.acquire();
                virtualBytes++;
                System.out.println(Thread.currentThread().getName() + " writes");
                roomEmpty.release();
            } catch (InterruptedException e) {
                System.out.println(e.getMessage());
            }
        }
    }
    public static void main(String[] args) throws Exception {
        
        CountDownLatch latch = new CountDownLatch(2000);
        
        long startTime = Instant.now().toEpochMilli();
        
        Reader read = new Reader();
        Writer write = new Writer();
        
        for(int i = 0; i <1000; i++) {
            Thread t10 = new Thread(write);
            Thread t11 = new Thread(read);
            t10.setName("writer " + Integer.toString(i));
            t11.setName("reader " + Integer.toString(i));
            t10.start();
            t11.start();
        }
        
        for(int i = 0; i <= 2000; i++){
            latch.countDown();
        }

        latch.await();
        Thread.sleep(100);
        long endTime = Instant.now().toEpochMilli();

        long timeElapsed = endTime - startTime;
        
        System.out.println(timeElapsed + "ms");
    }
}