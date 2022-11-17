import java.io.File;
import java.io.FileNotFoundException;
import java.util.Scanner;

public class main {

    static int[][] garden;

    public static void main(String[] args) throws FileNotFoundException {
        File input = new File("input.txt");
        Scanner scanner = new Scanner(input).useDelimiter("\r\n");

        garden = new int[1000][1000];
        while (scanner.hasNext()) {
            useLights(scanner.next(), 0);
        }
        int totalLightsOn = getLightsOn();

        garden = new int[1000][1000];
        scanner = new Scanner(input).useDelimiter("\r\n");
        while (scanner.hasNext()) {
            useLights(scanner.next(), 1);
        }
        int totalLightLevel = getTotalLightLevel();

        scanner.close();

        System.out.println("Day 6 - Part 1");
        System.out.printf("You have '%d' light bulbs on\n", totalLightsOn);
        System.out.println("Day 6 - part 2");
        System.out.printf("The total light level is: '%d'", totalLightLevel);
    }

    enum actions {
        off, on, toggle
    }

    public static void useLights(String input, int version) {
        String[] subStrings = input.split(" ");
        if (subStrings.length == 4) {
            turn(actions.toggle, subStrings[1], subStrings[3], version);
        }
        if (subStrings[1].equalsIgnoreCase("On")) {
            turn(actions.on, subStrings[2], subStrings[4], version);
        }
        if (subStrings[1].equalsIgnoreCase("Off")) {
            turn(actions.off, subStrings[2], subStrings[4], version);
        }
    }

    public static void turn(actions a, String startString, String endString, int version) {
        String[] start = startString.split(",");
        String[] end = endString.split(",");

        for (int x = Integer.parseInt(start[0]); x <= Integer.parseInt(end[0]); x++) {
            for (int y = Integer.parseInt(start[1]); y <= Integer.parseInt(end[1]); y++) {
                if (version == 0) {
                    actionOld(a, x, y);
                } else {
                    actionNew(a, x, y);
                }
            }
        }
    }

    public static void actionOld(actions a, int x, int y) {
        if (a == actions.on) {
            garden[x][y] = 1;
        }
        if (a == actions.off) {
            garden[x][y] = 0;
        }
        if (a == actions.toggle) {
            if (garden[x][y] == 0) {
                garden[x][y] = 1;
            } else {
                garden[x][y] = 0;
            }
        }
    }

    public static void actionNew(actions a, int x, int y) {
        if (a == actions.on) {
            garden[x][y] += 1;
        }
        if (a == actions.off) {
            if (garden[x][y] > 0)
                garden[x][y] -= 1;
        }
        if (a == actions.toggle) {
            garden[x][y] += 2;
        }
    }

    public static int getLightsOn() {
        int result = 0;
        for (int[] x : garden) {
            for (int y : x) {
                if (y == 1) result++;
            }
        }
        return result;
    }

    public static int getTotalLightLevel() {
        int result = 0;
        for (int[] x : garden) {
            for (int y : x) {
                result += y;
            }
        }
        return result;
    }
}
