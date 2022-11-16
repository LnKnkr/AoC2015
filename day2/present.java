public class present {
    int length;
    int width;
    int height;
    int totalPaper;

    public present(String input) {
        String[] metrics = input.split("x");
        length = Integer.parseInt(metrics[0]);
        width = Integer.parseInt(metrics[1]);
        height = Integer.parseInt(metrics[2]);
        calculateTotalPaper();
    }

    public void calculateTotalPaper() {
        int lengthWidth = length * width;
        int widthHeight = width * height;
        int heightLength = height * length;

        if (lengthWidth <= widthHeight && lengthWidth <= heightLength)
            totalPaper = lengthWidth;
        if (widthHeight <= lengthWidth && widthHeight <= heightLength)
            totalPaper = widthHeight;
        if (heightLength <= widthHeight && heightLength <= lengthWidth)
            totalPaper = heightLength;
        System.out.println(totalPaper);

        totalPaper += (2 * lengthWidth) + (2 * heightLength) + (2 * widthHeight);
    }
}
