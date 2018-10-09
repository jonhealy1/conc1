import java.util.concurrent.Semaphore;
//
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
                System.out.println("Thread "+Thread.currentThread().getName() + " is READING");
                Thread.sleep(1500);
                System.out.println("Thread "+Thread.currentThread().getName() + " has FINISHED READING");

                //Releasing section
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
                System.out.println("Thread "+Thread.currentThread().getName() + " is WRITING");
                Thread.sleep(2500);
                System.out.println("Thread "+Thread.currentThread().getName() + " has finished WRITING");
                roomEmpty.release();
            } catch (InterruptedException e) {
                System.out.println(e.getMessage());
            }
        }
    }
    public static void main(String[] args) throws Exception {
        Reader read = new Reader();
        Writer write = new Writer();
        
        Thread t1 = new Thread(read);
        t1.setName("thread1");
        Thread t2 = new Thread(read);
        t2.setName("thread2");
        Thread t3 = new Thread(write);
        t3.setName("thread3");
        Thread t4 = new Thread(read);
        t4.setName("thread4");
        t1.start();
        t3.start();
        t2.start();
        t4.start();
        
        /*
        for(int i = 0; i <= 3; i++){
            String p;
            p = "read";
            Thread t1 = new Thread(read);
            t1.setName(p);
        }
        for(int i = 0; i <= 3; i++){
            String p;
            p = "write";
            Thread t1 = new Thread(write);
            t1.setName(p);
        }*/
        Thread.sleep(12500);
    }
}