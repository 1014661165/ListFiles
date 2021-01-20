import bean.ListFileTask;
import conf.Config;
import org.apache.commons.io.FileUtils;
import java.io.File;
import java.io.IOException;
import java.time.Duration;
import java.util.List;
import java.util.concurrent.ForkJoinPool;

public class Main {

    public static void main(String[] args) {
        long start = System.currentTimeMillis();
        System.out.println("load config");
        File config = new File(Config.CONFIG_FILE);
        try {
            if (!config.exists()){
                System.out.printf("please update %s\n", Config.CONFIG_FILE);
                Config.save();
                System.exit(0);
            }
            Config.load();
        }catch (IOException e){
            e.printStackTrace();
            System.exit(2);
        }

        System.out.println("list files");
        ForkJoinPool pool = new ForkJoinPool();
        List<String> files = pool.invoke(new ListFileTask(Config.DATASET, Config.EXTENSION_MAP, true));

        System.out.printf("output file list to %s\n", Config.OUTPUT);
        File output = new File(Config.OUTPUT);
        output.getParentFile().mkdirs();
        try {
            FileUtils.writeLines(new File(Config.OUTPUT), files);
        }catch (IOException e){
            e.printStackTrace();
        }
        long end = System.currentTimeMillis();
        Duration duration = Duration.ofMillis(end - start);
        System.out.printf("task finish! time cost %s\n", duration.toString());
    }
}
