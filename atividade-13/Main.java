import java.util.Scanner;

class AFD {
    private String state;

    public AFD() {
        state = "q0";
    }

    public void transition(char input) {
        switch (state) {
            case "q0":
                if (input == '1') {
                    state = "q1"; 
                } else if (input == '0') {
                    state = "q-1"; 
                }
                break;
            case "q1":
                if (input == '1') {
                    state = "q2"; 
                } else if (input == '0') {
                    state = "q0"; 
                }
                break;
            case "q-1":
                if (input == '1') {
                    state = "q0"; 
                } else if (input == '0') {
                    state = "q-2"; 
                }
                break;
            case "q2":
                if (input == '1') {
                    state = "q3"; 
                } else if (input == '0') {
                    state = "q1"; 
                }
                break;
            case "q-2":
                if (input == '1') {
                    state = "q-1"; 
                } else if (input == '0') {
                    state = "q-3"; 
                }
                break;
        }
    }

    public boolean isAccepted(String input) {
        state = "q0"; 
        for (char c : input.toCharArray()) {
            transition(c);
        }

        return state.equals("q1") || state.equals("q2") || state.equals("q3");
    }

    
}

public class Main {

    public static void main(String[] args) {
        AFD afd = new AFD();
        System.out.println(afd.isAccepted("110"));  
        System.out.println(afd.isAccepted("101"));  
        System.out.println(afd.isAccepted("11100"));
        System.out.println(afd.isAccepted("100"));  
        System.out.println(afd.isAccepted("111"));  
    }
}