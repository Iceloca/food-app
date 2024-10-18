package com.example.foodapp.Adapter;

import android.annotation.SuppressLint;
import android.content.Context;
import android.content.Intent;
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
import com.example.foodapp.Activity.DetailActivity;
import com.example.foodapp.Domain.Product.ProductResponse;
import com.example.foodapp.R;

import org.jetbrains.annotations.NotNull;

import java.util.ArrayList;

public class ProductListAdapter extends RecyclerView.Adapter<ProductListAdapter.viewholder> {
    private ArrayList<ProductResponse> items;
    private Context context;

    public ProductListAdapter(ArrayList<ProductResponse> items) {
        this.items = items;
    }


    @NonNull
    @Override
    public ProductListAdapter.viewholder onCreateViewHolder(@NonNull ViewGroup parent, int viewType) {
        context = parent.getContext();
        View inflate = LayoutInflater
                .from(parent.getContext())
                .inflate(R.layout.view_holder_productlist, parent, false);
        return new viewholder(inflate);
    }

    @SuppressLint("SetTextI18n")
    @Override
    public void onBindViewHolder(@NonNull ProductListAdapter.viewholder holder, int position) {
        holder.nameProductList.setText(items.get(position).getName());
        holder.timeProductList.setText("20мин");
        holder.priceProductList.setText(items.get(position).getPrice() + "BYN");

        Glide.with(context)
                .load(items.get(position).getImageURL())
                .transform(new CenterCrop(), new RoundedCorners(30))
                .into(holder.imageProductList);

        holder.itemView.setOnClickListener(view -> {
            Intent intent = new Intent(context, DetailActivity.class);
            intent.putExtra("product", items.get(position));
            context.startActivity(intent);
        });
    }

    @Override
    public int getItemCount() {
        return items.size();
    }

    public static class viewholder extends RecyclerView.ViewHolder {
        TextView nameProductList, priceProductList, timeProductList;
        ImageView imageProductList;

        public viewholder(@NotNull View itemView) {
            super(itemView);
            nameProductList = itemView.findViewById(R.id.nameProductList);
            priceProductList = itemView.findViewById(R.id.priceProductList);
            timeProductList = itemView.findViewById(R.id.timeProductList);
            imageProductList = itemView.findViewById(R.id.imageProductList);
        }
    }
}
