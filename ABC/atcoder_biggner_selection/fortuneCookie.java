package code_interview.atcoder_biggner_selection;
import java.util.*;

public class fortuneCookie {
	public static void main(String[] args) {
		Scanner sc = new Scanner(System.in);
		int a = sc.nextInt();
		int b = sc.nextInt();
		int c = sc.nextInt();
		int d = sc.nextInt();

		if ((a==(b+c+d))   ||
			(b==(a+c+d))   ||
			(c==(a+b+d))   ||
			(d==(a+b+c))   ||
			((a+b)==(c+d)) ||
			((a+c)==(b+d)) ||
			((a+d)==(b+c))) {
			System.out.println("Yes");
		} else {
			System.out.println("No");
		}

	}

}