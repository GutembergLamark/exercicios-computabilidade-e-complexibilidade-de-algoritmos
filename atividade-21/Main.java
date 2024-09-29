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
        states.add("q_reject");
        startState = "q0";
        acceptStates = new HashSet<>();
        acceptStates.add("q0");
    }

    public boolean isAccepted(String input) {
        Set<String> currentStates = new HashSet<>();
        currentStates.add(startState);

        for (char c : input.toCharArray()) {
            Set<String> nextStates = new HashSet<>();
            for (String state : currentStates) {
                if (state.equals("q0")) {
                    if (c == 'a') {
                        nextStates.add("q0");
                    } else if (c == 'b') {
                        nextStates.add("q1");
                    }
                } else if (state.equals("q1")) {
                    if (c == 'a') {
                        nextStates.add("q0");
                    } else if (c == 'b') {
                        nextStates.add("q_reject");
                    }
                } else if (state.equals("q_reject")) {
                    nextStates.add("q_reject");
                }
            }
            currentStates = nextStates;
        }
        return currentStates.contains("q0");
    }

    
}

public class Main {

    public static void main(String[] args) {
        AFN afn = new AFN();
        System.out.println(afn.isAccepted("a"));       
        System.out.println(afn.isAccepted("ab"));      
        System.out.println(afn.isAccepted("abb"));     
        System.out.println(afn.isAccepted("ababa"));    
        System.out.println(afn.isAccepted("bba"));     
        System.out.println(afn.isAccepted("ba"));      
        System.out.println(afn.isAccepted("b"));       
    }
}