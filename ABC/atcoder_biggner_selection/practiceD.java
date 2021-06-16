package code_interview.atcoder_biggner_selection;
import java.util.Scanner;
import java.util.*;

public class practiceD {
	public static void main(String[] args) {
		Scanner sc = new Scanner(System.in);
		List<Integer> input = new ArrayList<>();

		int max = sc.nextInt();
		for (int i=0; i<max; i++) {
			input.add(sc.nextInt());
		}

		int maxCount = 1;
		for (Integer target : input) {
			int count = 0;
				//int o = 10;
			int o = countMax(target, count);
			if (maxCount >= o) {
				maxCount = o;
			}
		}


		System.out.println(maxCount);
	}

	private static int countMax(Integer target, int count) {
		System.out.println("target" + target);

		if (target <= 0) {
			return target;
		}
		if (target%2 == 0) {
			count = count + 1;
			System.out.println("count++ " + count);

			count = countMax(target/2, count);
		}

		System.out.println("count" + count);
		return count;
	}

}