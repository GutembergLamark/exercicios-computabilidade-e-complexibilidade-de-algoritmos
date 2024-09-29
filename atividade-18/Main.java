class AFD {
    private String state;

    public AFD() {
        state = "q00";  
    }

    public void transition(char input) {
        switch (state) {
            case "q00":
                if (input == '0') state = "q10";
                else if (input == '1') state = "q01";
                break;
            case "q01":
                if (input == '0') state = "q11";
                else if (input == '1') state = "q00";
                break;
            case "q10":
                if (input == '0') state = "q00";
                else if (input == '1') state = "q11";
                break;
            case "q11":
                if (input == '0') state = "q01";
                else if (input == '1') state = "q10";
                break;
        }
    }

    public boolean isAccepted(String input) {
        state = "q00";  
        for (char c : input.toCharArray()) {
            transition(c);
        }
        return state.equals("q11");  
    }

    
}

public class Main {

    public static void main(String[] args) {
        AFD afd = new AFD();
        System.out.println(afd.isAccepted("01"));    
        System.out.println(afd.isAccepted("1100"));  
        System.out.println(afd.isAccepted("10"));    
        System.out.println(afd.isAccepted("0000"));  
        System.out.println(afd.isAccepted("1111"));  
    }
}