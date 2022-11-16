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

        int[] out = calculateResults(list);
        System.out.printf("Day 2\nThe elves should buy '%d' foot wrapping paper\n" +
                "The elves will use a total of '%d' foot ribbon\n", out[0], out[1]);
    }

    public static int[] calculateResults(ArrayList<present> list) {
        int[] amount = new int[2];
        for (present p : list) {
            amount[0] += p.totalPaper;
            amount[1] += p.lengthRibbon;
        }
        return amount;
    }
}


