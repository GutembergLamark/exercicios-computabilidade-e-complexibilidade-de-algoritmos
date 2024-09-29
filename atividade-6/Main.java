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
        acceptStates.add("q1");
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
                    } else {
                        nextStates.add("q0");  
                    }
                } else if (state.equals("q1")) {
                    nextStates.add("q1");  
                }
            }
            currentStates = nextStates;
        }

        
        for (String state : currentStates) {
            if (acceptStates.contains(state)) {
                return true;
            }
        }
        return false;
    }
}

public class Main {

    public static void main(String[] args) {
        AFN afn = new AFN();
        System.out.println(afn.isAccepted("111"));  
        System.out.println(afn.isAccepted("110"));  
        System.out.println(afn.isAccepted("001"));  
        System.out.println(afn.isAccepted("1"));    
        System.out.println(afn.isAccepted("0"));    
    }
    
}