package bean;

import org.apache.commons.io.FilenameUtils;

import java.io.File;
import java.util.ArrayList;
import java.util.List;
import java.util.Map;
import java.util.concurrent.RecursiveTask;

public class ListFileTask extends RecursiveTask<List<String>> {

    private String path;
    private Map<String, Boolean> extensionMap;
    private boolean isRoot;

    public ListFileTask(String path, Map<String, Boolean> extensionMap, boolean isRoot) {
        this.path = path;
        this.extensionMap = extensionMap;
        this.isRoot = isRoot;
    }

    @Override
    protected List<String> compute() {
        List<String> paths = new ArrayList<>();
        File file = new File(path);
        File[] items = file.listFiles();
        if (items == null || items.length == 0){
            return paths;
        }
        int cnt = 0;
        for (File item: items){
            if (isRoot){
                cnt++;
                System.out.printf("%.2f%%\n", cnt*100f/items.length);
            }

            if (item.isFile()){
                String ext = FilenameUtils.getExtension(item.getName());
                if (extensionMap.containsKey(ext)){
                    paths.add(item.getAbsolutePath());
                }
            }else{
                ListFileTask task = new ListFileTask(item.getAbsolutePath(), extensionMap, false);
                task.fork();
                paths.addAll(task.join());
            }
        }
        return paths;
    }
}
