class AFD {
    private String state;

    public AFD() {
        state = "q0"; 
    }

    private void transition(char c) {
        switch (state) {
            case "q0":
                if (c == 'a') state = "q1";
                else if (c == 'b') state = "q2";
                break;
            case "q1":
                if (c == 'a') state = "q2";
                else if (c == 'b') state = "q0";
                break;
            case "q2":
                if (c == 'a') state = "q0";
                else if (c == 'b') state = "q1";
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
        System.out.println(afd.isAccepted("aaabbb"));  
        System.out.println(afd.isAccepted("aab"));     
        System.out.println(afd.isAccepted("aabb"));    
        System.out.println(afd.isAccepted("aaa"));     
        System.out.println(afd.isAccepted("bbb"));     
        System.out.println(afd.isAccepted("ababab"));  
    }
}