import java.util.Scanner;

public class MathCalculator {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        
        System.out.println("Enter first number:");
        int num1 = scanner.nextInt();
        
        System.out.println("Enter operator (+, -, *, /):");
        char op = scanner.next().charAt(0);
        
        System.out.println("Enter second number:");
        int num2 = scanner.nextInt();
        
        switch (op) {
            case '+':
                System.out.println(num1 + " " + op + " " + num2 + " = " + (num1 + num2));
                break;
            case '-':
                System.out.println(num1 + " " + op + " " + num2 + " = " + (num1 - num2));
                break;
            case '*':
                System.out.println(num1 + " " + op + " " + num2 + " = " + (num1 * num2));
                break;
            case '/':
                if (num2 != 0) {
                    System.out.println(num1 + " " + op + " " + num2 + " = " + (num1 / num2));
                } else {
                    System.out.println("Error: Division by zero is not allowed.");
                }
                break;
            default:
                System.out.println("Invalid operator. Please enter a valid operator (+, -, *, /).");
        }
        
        scanner.close();
    }
}
