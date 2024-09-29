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
        acceptStates.add("q1");
    }

    public boolean isAccepted(String input) {
        Set<String> currentStates = new HashSet<>();
        currentStates.add(startState);

        for (char c : input.toCharArray()) {
            Set<String> nextStates = new HashSet<>();
            for (String state : currentStates) {
                switch (state) {
                    case "q0":
                        if (c == 'a') {
                            nextStates.add("q0");
                        } else if (c == 'b') {
                            nextStates.add("q1");
                        }
                        break;
                    case "q1":
                        if (c == 'b') {
                            nextStates.add("q1");
                        }
                        break;
                }
            }
            currentStates = nextStates;
        }

        return currentStates.contains("q0") || currentStates.contains("q1");
    }
}

public class Main {

   public static void main(String[] args) {
        AFN afn = new AFN();
        System.out.println(afn.isAccepted("aaa"));    
        System.out.println(afn.isAccepted("aabbb"));  
        System.out.println(afn.isAccepted("aaab"));   
        System.out.println(afn.isAccepted("aabb"));   
        System.out.println(afn.isAccepted("abba"));   
        System.out.println(afn.isAccepted("b"));      
        System.out.println(afn.isAccepted("ba"));     
    }
}