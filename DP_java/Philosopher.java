//package com.baeldung.concurrent.diningphilosophers;

public class Philosopher implements Runnable {

    private final Object leftFork;
    private final Object rightFork;

    Philosopher(Object left, Object right) {
        this.leftFork = left;
        this.rightFork = right;
    }

    private void doAction(String action) throws InterruptedException {
        System.out.println(Thread.currentThread().getName() + " " + action);
        //Thread.sleep(((int) (Math.random() * 100)));
    }

    @Override
    public void run() {
        try {
            //while (true) {
            //for(int i = 0; i < 10; i++){
                doAction("thinking"); // thinking
                synchronized (leftFork) {
                    doAction("picked up left fork");
                    synchronized (rightFork) {
                        doAction("picked up right fork - eating"); // eating
                        doAction("put down right fork");
                    }
                    doAction("put down left fork. Full");
                }
           // }
        } catch (InterruptedException e) {
            Thread.currentThread().interrupt();
        }
    }
}