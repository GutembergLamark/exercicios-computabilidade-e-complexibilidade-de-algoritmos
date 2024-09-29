class AFD {
    private String state;

    public AFD() {
        this.state = "q0";
    }

    public void transition(char c) {
        if ("q0".equals(state)) {
            if (c == '0') {
                state = "q1";
            }
        }
    }

    public boolean isAccepted(String input) {
        state = "q0";  
        for (char c : input.toCharArray()) {
            transition(c);
        }
        return "q1".equals(state);  
    }
}

public class Main {

    public static void main(String[] args) {
        AFD afd = new AFD();
        System.out.println(afd.isAccepted("1"));     
        System.out.println(afd.isAccepted("0"));     
        System.out.println(afd.isAccepted("111"));   
        System.out.println(afd.isAccepted("101"));   
        System.out.println(afd.isAccepted("1100"));  
    }
    
}