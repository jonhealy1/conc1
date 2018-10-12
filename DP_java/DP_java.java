//package com.DPbaeldung.concurrent.diningphilosophers;
//https://github.com/eugenp/tutorials/blob/master/core-java-concurrency/src/main/java/com/baeldung/concurrent/diningphilosophers/Philosopher.java
import java.time.Instant;
import java.util.concurrent.*;

public class DP_java {

    public static void main(String[] args) throws Exception {

        long startTime = Instant.now().toEpochMilli();

        Philosopher[] philosophers = new Philosopher[5];
        Object[] forks = new Object[philosophers.length];

        for (int i = 0; i < forks.length; i++) {
            forks[i] = new Object();
        }
        for (int j = 0; j < 200; j++) {
        for (int i = 0; i < philosophers.length; i++) {

            Object leftFork = forks[i];
            Object rightFork = forks[(i + 1) % forks.length];

            if (i == philosophers.length - 1) {
                philosophers[i] = new Philosopher(rightFork, leftFork); // The last philosopher picks up the right fork first
            } else {
                philosophers[i] = new Philosopher(leftFork, rightFork);
            }

            Thread t = new Thread(philosophers[i], "phil " + (i + 1));
            t.start();
        }
        }
        Thread.sleep(100);
        long endTime = Instant.now().toEpochMilli();

        long timeElapsed = endTime - startTime;
        
        System.out.println(timeElapsed + "ms");
    }
}