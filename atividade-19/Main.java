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
        states.add("q3");

        startState = "q0";
        acceptStates = new HashSet<>();
        acceptStates.add("q3");
    }

    public boolean isAccepted(String input) {
        Set<String> currentStates = new HashSet<>();
        currentStates.add(startState);

        for (char c : input.toCharArray()) {
            Set<String> nextStates = new HashSet<>();
            for (String state : currentStates) {
                switch (state) {
                    case "q0":
                        if (c == '0') {
                            nextStates.add("q1");
                        }
                        break;
                    case "q1":
                        if (c == '1') {
                            nextStates.add("q2");
                        } else if (c == '0') {
                            nextStates.add("q1");
                        }
                        break;
                    case "q2":
                        if (c == '0') {
                            nextStates.add("q3");
                        } else if (c == '1') {
                            nextStates.add("q0");
                        }
                        break;
                    case "q3":
                        nextStates.add("q3");
                        break;
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
        System.out.println(afn.isAccepted("010"));   
        System.out.println(afn.isAccepted("0010"));  
        System.out.println(afn.isAccepted("111"));   
        System.out.println(afn.isAccepted("0110010"));  
        System.out.println(afn.isAccepted("101"));   
    }
}