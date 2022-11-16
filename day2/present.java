import java.util.ArrayList;

public class present {

    ArrayList<Integer> dimensions;
    int totalPaper;
    int lengthRibbon;

    public present(String input) {
        String[] metrics = input.split("x");
        dimensions = new ArrayList<>();
        dimensions.add(Integer.parseInt(metrics[0]));
        dimensions.add(Integer.parseInt(metrics[1]));
        dimensions.add(Integer.parseInt(metrics[2]));
        dimensions.sort(null);
        calculateTotalPaper();
        calculateRibbonLength();
    }

    public void calculateTotalPaper() {
        totalPaper = 3 * (dimensions.get(0) * dimensions.get(1));
        totalPaper += 2 * (dimensions.get(0) * dimensions.get(2));
        totalPaper += 2 * (dimensions.get(1) * dimensions.get(2));
    }

    public void calculateRibbonLength(){
        lengthRibbon = 2*dimensions.get(0) + 2* dimensions.get(1);
        lengthRibbon += dimensions.get(0) * dimensions.get(1) * dimensions.get(2);
    }
}
