package code_interview.atcoder_biggner_selection;

import java.util.*;
public class practiceC {
	public static void main(String[] args) {
		Scanner sc = new Scanner(System.in);
		String in = sc.next();

		int count = 0;

		if (in.contains("1")) {
			char[] arr = in.toCharArray();
			for (char c : arr) {
				if (c == '1') {
					count++;
				}
			}
			
		}

		System.out.println(count);
	}

}