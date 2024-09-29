class AFD {
    private String state;

    public AFD() {
        this.state = "q0";
    }

    private void transition(char c) {
        switch (state) {
            case "q0":
                if (c == '0') state = "q1";
                else if (c == '1') state = "q0";
                break;
            case "q1":
                if (c == '0') state = "q0";
                else if (c == '1') state = "q1";
                break;
        }
    }

    public boolean isAccepted(String input) {
        state = "q0"; 
        for (char c : input.toCharArray()) {
            transition(c);
        }
        return state.equals("q0");
    }

    
}

public class Main {

    public static void main(String[] args) {
        AFD afd = new AFD();
        System.out.println(afd.isAccepted("110"));  
        System.out.println(afd.isAccepted("101"));  
        System.out.println(afd.isAccepted("1001"));
        System.out.println(afd.isAccepted("0"));    
        System.out.println(afd.isAccepted("00"));   
    }
    
}