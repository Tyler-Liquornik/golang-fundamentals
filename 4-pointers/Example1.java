public class Example1 {
    public static void main(String[] args) {
        int x = 1;
        int y = x; // copy

        y = 100;

        System.out.println("x: " + x); // x == 1 (unchanged)
        System.out.println("y: " + y); // y == 100
    }
}
