import java.util.regex.Matcher;
import java.util.regex.Pattern;

class AFN {
    public static boolean isAccepted(String string) {
        Pattern pattern = Pattern.compile("110");
        Matcher matcher = pattern.matcher(string);
        return matcher.find();
    }
}

public class Main {

    public static void main(String[] args) {
        AFN afn = new AFN();
        System.out.println(afn.isAccepted("110"));  
        System.out.println(afn.isAccepted("101")); 
        System.out.println(afn.isAccepted("1110")); 
        System.out.println(afn.isAccepted("00110"));
        System.out.println(afn.isAccepted("000"));  
    }
    
}