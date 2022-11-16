import java.io.File;
import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.Scanner;

public class main {

    public static void main(String[] args) throws FileNotFoundException {
        Scanner scanner = new Scanner(new File("input.txt"));
        ArrayList<present> list = new ArrayList<>();
        while (scanner.hasNext()) {
            list.add(new present(scanner.next()));
        }
        System.out.printf("Day 2 - Part 1\nThe elves should buy '%d' foot wrapping paper\n", totalAmountWrappingPaper(list));
    }

    public static int totalAmountWrappingPaper(ArrayList<present> list) {
        int amount = 0;
        for (present p : list) {
            amount += p.totalPaper;
        }
        return amount;
    }
}


