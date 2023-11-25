import java.util.Scanner;

public class Main {
  public static void main(String[] args){
    // You can use print statements as follows for debugging, they'll be visible when running tests.
    System.out.println("Logs from your program will appear here!");

    if(args.length < 2 || !args[0].equals("-E")) {
      System.out.println("Expected first argument to be -E");
      return;
    }

    final String pattern = args[1];
    final Scanner input = new Scanner(System.in);
    final String inputLine = input.nextLine();

    // Uncomment this block to pass the first stage
    // if (inputLine.contains(pattern)) {
    //   System.exit(0);
    // } else {
    //   System.exit(1);
    // }
  }
}
