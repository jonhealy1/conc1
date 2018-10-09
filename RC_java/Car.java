//============================ CAR CLASS ===========================
class Car implements Runnable {
    private int id; // Car ID
    private Monitor carMon;
    public Car(int i, Monitor monitorIn) {
      id = i;
      this.carMon = monitorIn;
    }
    public void run() {
      while(true) {
        carMon.passengerGetOn(id);
        try{
          Thread.sleep((int)(Math.random()*2000));
        }catch(InterruptedException e){
        } // Car runs for a while
        carMon.passengerGetOff(id);
      }
    }
  } // end of Car class