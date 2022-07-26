package offer;

class Solution2 {
    public TreeNode lowestCommonAncestor(TreeNode root,TreeNode p,TreeNode q){
        if (root == null || p == root || q == root){
            return root;
        }
        TreeNode left = lowestCommonAncestor(root.left,p,q);
        TreeNode right = lowestCommonAncestor(root.right,p,q);
        if(right == null){
            return left;
        }
        if(left == null){
            return right;
        }
        return root;
    }

    public static void main(String[] args) {

    }
}
