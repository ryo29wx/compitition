package code_interview.atcoder_biggner_selection;
import java.util.Scanner;

public class practiceB {
	public static void main(String[] args) {
		Scanner sc = new Scanner(System.in);
		int a = sc.nextInt();
		int b = sc.nextInt();

		if (a*b%2 == 0) {
			System.out.println("Even");
		} else {
			System.out.println("Odd");
		}

	}
}