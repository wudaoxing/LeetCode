/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode(int x) { val = x; }
 * }
 */
public class Codec {

    // Encodes a tree to a single string.
    public String serialize(TreeNode root) {
        if(root == null){
            return "[]";
        }
        StringBuilder res = new StringBuilder("[");
        Queue<TreeNode> queue = new LinkedList<TreeNode>();
        queue.add(root);
        while(queue.size() != 0){
            TreeNode node = queue.remove();
            if(node != null){
                 res.append(node.val+",");
                 queue.add(node.left);
                 queue.add(node.right);
            }else{
                res.append("null,");
            }
        }
        res.deleteCharAt(res.length()-1);
        res.append("]");
        return res.toString();
    }

    // Decodes your encoded data to tree.
    public TreeNode deserialize(String data) {
        if(data.equals("[]")){
            return null;
        }
        String[] strs = data.substring(1,data.length()-1).split(",");
        TreeNode root = new TreeNode(Integer.parseInt(strs[0]));
        Queue<TreeNode> queue = new LinkedList<TreeNode>(){{add(root);}};
        int i = 1;
        while(!queue.isEmpty()){
            TreeNode node = queue.poll();
            if(!strs[i].equals("null")){
                node.left = new TreeNode(Integer.parseInt(strs[i]));
                queue.add(node.left);
            }
            i++;
            if(!strs[i].equals("null")){
                node.right = new TreeNode(Integer.parseInt(strs[i]));
                queue.add(node.right);
            }
            i++;
        }
        return root;
    }


}

// Your Codec object will be instantiated and called as such:
// Codec codec = new Codec();
// codec.deserialize(codec.serialize(root));