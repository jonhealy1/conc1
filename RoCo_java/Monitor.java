//=========================== Monitor Class ================================
class Monitor {
    private int i, line_length; // Number of passengers waiting to board the car.
    private int seats_available = 0;
    boolean coaster_loading_passengers = false;
    boolean passengers_riding = true;
  
    private Object notifyPassenger = new Object(); // enter/exit protocol provides mutual exclusion.
    private Object notifyCar = new Object(); // the car waits on this.
  
    public void tryToGetOnCar(int i) {
      synchronized (notifyPassenger) {
        while (!seatAvailable()) {
          try {
            notifyPassenger.wait(); // Notify the passenger to wait
            } catch (InterruptedException e){}
        }
      }
      System.out.println("Passenger "+ i + " gets in car at timestampe: " + System.currentTimeMillis());
      synchronized (notifyCar) {notifyCar.notify();}
    }
  
    private synchronized boolean seatAvailable() {
      // Check if seat is still available for passenger who tries to get on.
      if ((seats_available > 0)
          && (seats_available <= RollerCoaster.SEAT_AVAIL)
          && (!passengers_riding)) {
        seats_available--;
        return true;
      } else return false;
    }
  
    public void passengerGetOn(int i) {
      synchronized (notifyCar) {
        while (!carIsRunning()) {
          try {
            notifyCar.wait();
            } catch (InterruptedException e){}
        }
      }
      System.out.println("The Car is full and starts running at timestampe: "+ System.currentTimeMillis());
      synchronized(notifyPassenger) {notifyPassenger.notifyAll();}
    }
  
    private synchronized boolean carIsRunning() {
      // Check if car is running
      if (seats_available == 0) {
        //if there is no seat, car starts to run and reset parameters.
        seats_available = RollerCoaster.SEAT_AVAIL;
        // reset seat available num for the next ride
        coaster_loading_passengers = true; // Indicating car is running.
        passengers_riding = true; // passengers are riding in the car.
        return true;
      } else return false;
    }
  
    public void passengerGetOff(int i) {
      synchronized (this) {
        // reset parameters
        passengers_riding = false;
        coaster_loading_passengers = false;
      }
      synchronized(notifyPassenger) {notifyPassenger.notifyAll();}
    }
  } // end of Monitor class