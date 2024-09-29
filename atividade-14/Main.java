import java.util.HashSet;
import java.util.Set;

class AFN {
    private Set<String> states;
    private String startState;
    private Set<String> acceptStates;

    public AFN() {
        states = new HashSet<>();
        states.add("q0");
        states.add("q1");
        states.add("q2");
        startState = "q0";
        acceptStates = new HashSet<>();
        acceptStates.add("q0");
        acceptStates.add("q1");
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
                    } else if (c == '1') {
                        nextStates.add("q2");
                    }
                } else if (state.equals("q1") && c == '1') {
                    nextStates.add("q2");
                } else if (state.equals("q2") && c == '0') {
                    nextStates.add("q1");
                }
            }
            currentStates = nextStates;
        }

        return !currentStates.isEmpty();
    }
}

public class Main {

    public static void main(String[] args) {
        AFN afn = new AFN();
        System.out.println(afn.isAccepted("010101"));  
        System.out.println(afn.isAccepted("101010"));  
        System.out.println(afn.isAccepted("1101"));    
        System.out.println(afn.isAccepted("100"));     
        System.out.println(afn.isAccepted(""));        
    }
}