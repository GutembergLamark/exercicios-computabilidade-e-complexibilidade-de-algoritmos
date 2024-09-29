class AFD {
    private int state;

    public AFD() {
        state = 0;
    }

    public boolean isAccepted(String string) {
        state = 0;
        for (char c : string.toCharArray()) {
            if (state == 0) {
                if (c == '1') {
                    state = 0;
                } else if (c == '0') {
                    state = 1;
                }
            } else if (state == 1) {
                if (c == '1') {
                    state = 0;
                } else if (c == '0') {
                    state = 2;
                }
            } else if (state == 2) {
                if (c == '0') {
                    state = 2;
                } else if (c == '1') {
                    
                }
            }
        }
        return state == 2;
    }
}

public class Main {

    public static void main(String[] args) {
        AFD afd = new AFD();
        System.out.println(afd.isAccepted("100011"));
        System.out.println(afd.isAccepted("10101"));
        System.out.println(afd.isAccepted("000"));
        System.out.println(afd.isAccepted("000110"));
        System.out.println(afd.isAccepted("1"));
        System.out.println(afd.isAccepted("100"));
    }
}