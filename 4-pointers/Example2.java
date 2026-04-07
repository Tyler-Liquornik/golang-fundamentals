public class Example2 {
    public static void process(int x) {
        x = 100;
    }

    public static void main(String[] args) {
        int x = 1;
        process(x);

        System.out.println("x after process(x): " + x); // x == 1 (unchanged)
    }
}
