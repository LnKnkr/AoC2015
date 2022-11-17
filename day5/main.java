import java.io.File;
import java.io.FileNotFoundException;
import java.util.Scanner;

public class main {
    public static void main(String[] args) throws FileNotFoundException {
        File text = new File("input.txt");
        Scanner scanner = new Scanner(text);

        int amountOldWays = 0;
        while (scanner.hasNext()) {
            if (isNiceStringOld(scanner.next())) {
                amountOldWays++;
            }
        }

        scanner = new Scanner(text);
        int amountNewWays = 0;
        while (scanner.hasNext()) {
            if (isNiceStringNew(scanner.next())) {
                amountNewWays++;
            }
        }

        scanner.close();

        System.out.println("Day 5 - Part 1");
        System.out.printf("Santa has '%d' nice strings\n", amountOldWays);

        System.out.println("Day 5 - Part 2");
        System.out.printf("Santa's new ways bring him '%d' nice strings", amountNewWays);
    }

    public static boolean isNiceStringOld(String input) {
        if (input.contains("ab") || input.contains("cd") || input.contains("pq") || input.contains("xy")) {
            return false;
        }
        char[] inputChars = input.toCharArray();
        int amountVowels = 0;
        boolean hasDoubleLetter = false;
        for (int i = 0; i < inputChars.length; i++) {
            if (inputChars[i] == 'a' || inputChars[i] == 'e' || inputChars[i] == 'i' || inputChars[i] == 'o' || inputChars[i] == 'u')
                amountVowels++;
            if (i + 1 < inputChars.length && inputChars[i] == inputChars[i + 1])
                hasDoubleLetter = true;
        }
        return hasDoubleLetter && amountVowels >= 3;
    }

    public static boolean isNiceStringNew(String input) {
        char[] inputChars = input.toCharArray();
        boolean hasSameLetterWithOneOffset = false;
        boolean hasPair = false;

        for (int i = 0; i < inputChars.length; i++) {
            if (i + 2 < inputChars.length && inputChars[i] == inputChars[i + 2])
                hasSameLetterWithOneOffset = true;
            for (int c = i + 2; c < inputChars.length; c++) {
                if (c + 1 < inputChars.length && inputChars[c] == inputChars[i] && inputChars[c + 1] == inputChars[i + 1]) {
                    hasPair = true;
                    break;
                }
            }
        }
        if (hasPair && hasSameLetterWithOneOffset)
            System.out.println(input);
        return hasPair && hasSameLetterWithOneOffset;
    }
}
