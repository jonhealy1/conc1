//========================= PASSENGER CLASS ===========================
class Passenger implements Runnable {
    private int id;
    private Monitor passengerMon;
    public Passenger(int i, Monitor monitorIn) {
      id = i;
      this.passengerMon = monitorIn;
    }
    public void run() {
      //while(true) {
      for(int i=0; i<1; i++){ // every passenger only takes one ride
        try{
          Thread.sleep((int)( Math.random()*2000));
        }catch(InterruptedException e){
        }
        passengerMon.tryToGetOnCar(id);
      }
    }
  } // end of Passenger class