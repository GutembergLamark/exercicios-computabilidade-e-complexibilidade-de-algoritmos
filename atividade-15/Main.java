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
                    nextStates.add("q1");
                } else if (state.equals("q1")) {
                    nextStates.add("q0");
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
        System.out.println(afn.isAccepted("ab"));    
        System.out.println(afn.isAccepted("abab"));  
        System.out.println(afn.isAccepted("a"));     
        System.out.println(afn.isAccepted("abc"));   
    }
}