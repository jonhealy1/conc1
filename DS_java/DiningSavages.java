import java.util.Random;
import java.util.concurrent.Semaphore;

//https://github.com/natashadecoste/ConcurrencyStudies/blob/master/DiningSavages.java

public class DiningSavages {
	public Random randomElement;
	public static Semaphore accessToFood;
	public int servingsAvailable;
	public int servings;

	// Savages
	class Savage implements Runnable {
		String name;

		public Savage(String n) {
			name = n;

		}

		public void begin() {
			(new Thread(new Savage(name))).start();
		}

		public void run() {
			while (true) {
				// try to take a portion
				while(servingsAvailable > 0){
				try {
					Thread.sleep(randomElement.nextInt(7) * 1000);
					accessToFood.acquire();

				} catch (InterruptedException e) {
					e.printStackTrace();
				}
				getServing(name);

				}
			}
		}

		private synchronized void getServing(String name) {
			if(servingsAvailable > 0){
				System.out.println("Savage " + name + " getting serving");
				servingsAvailable --;
				
			}else {
				System.out.println("Savage " + name + " checked the pot and theres nothing there!");
			}

			accessToFood.release();
		}

	}

	// Cook
	public class Cook implements Runnable {
		
		public Cook(){
			System.out.println("Cook created");
		}
		
		public void begin(){
			(new Thread(new Cook())).start();
		}
		
		public void run(){
			System.out.println("running");
			
			while(true){
				refillPot();

				
			}
		}
		
		public void refillPot(){
			if(servingsAvailable < 1){
				System.out.println("refilling");
	
				servingsAvailable = servings;
			}



		}

		
	} // end of Cook Class
	
	
	

	
	public DiningSavages(int numOfSavages, int servingSize){
		servings = servingSize;
		accessToFood = new Semaphore(1);
		servingsAvailable = servings;
		
		randomElement = new Random();
		
		Cook chef = new Cook();
		chef.begin();
		
		for(int i=0; i< numOfSavages; i++){
			Savage t = new Savage(Integer.toString(i));
			t.begin();
			
		}
		
	}
	
	
	
	
	
	public static void main(String[] args) {
		DiningSavages s = new DiningSavages(4,6);
	}

}