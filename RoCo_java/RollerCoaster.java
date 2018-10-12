import java.math.*;

//https://www.cs.mcgill.ca/~ylin30//courses/cs623/RollerCoaster/RollerCoaster.java

class RollerCoaster {
  public static int PASSENGER_NUM = 25;// Number of people totally
  public static int CAR_NUM = 1; // Number of passenger cars
  public static int SEAT_AVAIL = 4; // Number of passengers a car holds
  public static void main(String[] args) {
    Monitor rcMon = new Monitor();

    Car theCar;
    Passenger aPassenger;

    /* Create arrays of threads for initialization */
    Thread t1[] = new Thread[PASSENGER_NUM];
    Thread t2[] = new Thread[CAR_NUM];
    /* Fill the thread arrays */
    for (int i = 0; i < PASSENGER_NUM; i++) {
      aPassenger = new Passenger(i, rcMon);
      t1[i] = new Thread(aPassenger);
    }
    for (int i = 0; i < CAR_NUM; i++) {
      theCar = new Car(i, rcMon);
      t2[i] = new Thread(theCar);
    }

    for(int i = 0; i < PASSENGER_NUM; i++) {
      t1[i].start();
    }
    for(int i = 0; i < CAR_NUM; i++) {
      t2[i].start();
    }

    try {
      for (int i = 0; i < PASSENGER_NUM; i++) {
        t1[i].join();
      }
    } catch (InterruptedException e) {
      System.err.println("Passenger thread join interruption");
    }

    try {
      for (int i = 0; i < CAR_NUM; i++) {
        t2[i].join();
      }
    } catch (InterruptedException e) {
      System.err.println("Car thread join interruption");
    }
    System.out.println("Program has terminated Normally");
  }
} // end of RollerCoaster