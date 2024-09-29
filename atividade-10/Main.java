import java.util.HashSet;
import java.util.Set;

class AFN {
    private final Set<String> states = new HashSet<>();
    private final String startState = "q0";
    private final Set<String> acceptStates = new HashSet<>();

    public AFN() {
        states.add("q0");
        states.add("q1");
        states.add("q2");
        acceptStates.add("q2");
    }

    public boolean isAccepted(String input) {
        Set<String> currentStates = new HashSet<>();
        currentStates.add(startState); 

        for (char c : input.toCharArray()) {
            Set<String> nextStates = new HashSet<>();
            for (String state : currentStates) {
                if (state.equals("q0")) {
                    if (c == '0') {
                        nextStates.add("q1");
                    }
                } else if (state.equals("q1")) {
                    if (c == '1') {
                        nextStates.add("q2");
                    } else if (c == '0') {
                        nextStates.add("q1");
                    }
                } else if (state.equals("q2")) {
                    nextStates.add("q2");
                }
            }
            currentStates = nextStates;
        }

        return currentStates.stream().anyMatch(acceptStates::contains);
    }
}

public class Main {

    public static void main(String[] args) {
        AFN afn = new AFN();
        System.out.println(afn.isAccepted("010"));  
        System.out.println(afn.isAccepted("0011")); 
        System.out.println(afn.isAccepted("111"));  
        System.out.println(afn.isAccepted("000"));  
        System.out.println(afn.isAccepted("000111"));
    }
    
}