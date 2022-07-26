package offer;

import java.util.Scanner;

class TreeNode{
    int val;
    TreeNode left;
    TreeNode right;
    TreeNode(int x){val = x;}
}
class Solution{
    public TreeNode lowestCommonAncestor(TreeNode root,TreeNode p,TreeNode q){
        while (root != null){
            if (p.val < root.val && q.val < root.val){
                root = root.left;
            }else if(p.val > root.val && q.val > root.val){
                root = root.right;
            }else {
                break;
            }
        }
        return root;
    }

    public TreeNode buildTree(String[] arr, int index){
        if(arr[index].equals("null")){
            return null;
        }
        TreeNode root = new TreeNode(Integer.parseInt(arr[index]));
        int lnode = 2 * index + 1;
        int rnode = 2 * index + 2;
        if (lnode < arr.length){
            root.left = buildTree(arr,lnode);
        }
        if (rnode < arr.length){
            root.right = buildTree(arr,rnode);
        }
        return root;
    }
    public static void main(String[] args) {
        Solution s = new Solution();
        Scanner sca = new Scanner(System.in);
        String strs = sca.nextLine();
        String[] arr = strs.split(",");
        System.out.println("p=");
        TreeNode p = new TreeNode(sca.nextInt());
        System.out.println("q=");
        TreeNode q = new TreeNode(sca.nextInt());
        TreeNode root = s.buildTree(arr,0);
        System.out.println(s.lowestCommonAncestor(root,p,q).val);

    }
}