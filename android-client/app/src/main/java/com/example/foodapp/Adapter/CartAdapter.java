package com.example.foodapp.Adapter;

import android.annotation.SuppressLint;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.ImageView;
import android.widget.TextView;

import androidx.annotation.NonNull;
import androidx.recyclerview.widget.RecyclerView;

import com.bumptech.glide.Glide;
import com.bumptech.glide.load.resource.bitmap.CenterCrop;
import com.bumptech.glide.load.resource.bitmap.RoundedCorners;
import com.example.foodapp.Domain.Cart.Item;
import com.example.foodapp.Domain.Product.ProductResponse;
import com.example.foodapp.Management.CartManagement;
import com.example.foodapp.R;

import java.util.ArrayList;

public class CartAdapter extends RecyclerView.Adapter<CartAdapter.viewholder> {

    private final ArrayList<Item> items;
    private final CartManagement cartManagement;
    private final String cartId;

    public CartAdapter(ArrayList<Item> items, CartManagement cartManagement, String cartId) {
        this.items = items;
        this.cartManagement = cartManagement;
        this.cartId = cartId;
    }

    @NonNull
    @Override
    public CartAdapter.viewholder onCreateViewHolder(@NonNull ViewGroup parent, int viewType) {
        View inflate = LayoutInflater.from(parent.getContext()).inflate(R.layout.viewholder_cart, parent, false);
        return new viewholder(inflate);
    }

    @SuppressLint("SetTextI18n")
    @Override
    public void onBindViewHolder(@NonNull CartAdapter.viewholder holder, @SuppressLint("RecyclerView") int position) {
        ProductResponse product = items.get(position).getProduct();
        float price  = product.getPrice();
        int count = items.get(position).getCount();
        float totalPrice = price * count;

        holder.productName.setText(product.getName());
        holder.itemPrice.setText(price + "BYN");
        holder.countItems.setText(String.valueOf(count));
        holder.totalPrice.setText(totalPrice + "BYN");

        Glide.with(holder.itemView.getContext())
                .load(product.getImageURL())
                .transform(new CenterCrop(), new RoundedCorners(30))
                .into(holder.imageProduct);

        holder.plusItem.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View view) {
                cartManagement.addItem(items.get(position).getProduct(), cartId).thenAccept(success -> {
                    int newCount = count + 1;
                    items.get(position).setCount(newCount);
                    holder.countItems.setText(String.valueOf(newCount));
                    notifyItemChanged(position);
                });
            }
        });

        holder.minusItem.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View view) {
                cartManagement.deleteItem(cartId, items.get(position).getProduct().getId()).thenAccept(success -> {
                    int newCount = count - 1;
                    if (newCount == 0) {
                        items.remove(position);
                    } else {
                        items.get(position).setCount(newCount);
                    }
                    notifyItemChanged(position);
                });
            }
        });
    }

    @Override
    public int getItemCount() {
        return this.items.size();
    }

    public class viewholder extends RecyclerView.ViewHolder {
        TextView productName;
        TextView itemPrice, totalPrice;
        TextView countItems;
        ImageView imageProduct;
        TextView plusItem, minusItem;

        public viewholder(@NonNull View itemView) {
            super(itemView);

            productName = itemView.findViewById(R.id.nameProductCart);
            itemPrice = itemView.findViewById(R.id.priceItemCart);
            totalPrice = itemView.findViewById(R.id.totalPriceItems);
            countItems = itemView.findViewById(R.id.countItemsCart);
            imageProduct = itemView.findViewById(R.id.imageCart);
            plusItem = itemView.findViewById(R.id.plusItemCart);
            minusItem = itemView.findViewById(R.id.minusitemCart);
        }
    }
}
