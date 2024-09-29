class AFD {
    private String state;

    public AFD() {
        this.state = "q0"; 
    }

    public void transition(char c) {
        if (state.equals("q0")) {
            if (c == 'a') {
                state = "q1"; 
            }
        } else if (state.equals("q1")) {
            if (c == 'a') {
                state = "q0"; 
            }
        }
    }

    public boolean isAccepted(String input) {
        state = "q0"; 
        for (char c : input.toCharArray()) {
            transition(c);
        }
        return state.equals("q1"); 
    }
}

public class Main {

    public static void main(String[] args) {
        AFD afd = new AFD();
        System.out.println(afd.isAccepted("a"));  
        System.out.println(afd.isAccepted("aa"));
        System.out.println(afd.isAccepted("ab")); 
        System.out.println(afd.isAccepted("bba"));
        System.out.println(afd.isAccepted("bb")); 
    }
    
}