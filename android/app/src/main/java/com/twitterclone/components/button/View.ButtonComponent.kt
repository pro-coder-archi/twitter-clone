package com.twitterclone.components.button

import androidx.compose.foundation.Image
import androidx.compose.foundation.layout.*
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material.Button
import androidx.compose.material.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import com.twitterclone.theme.AppTheme

@Composable
fun ButtonComponentView(

    buttonText: String,
    clickHandler: ( ) -> Unit,
    leadingDrawableResourceID: Int?= null,
    trailingDrawableResourceID: Int?= null,
    variant: ButtonComponentVariants= ButtonComponentVariants.DEFAULT,
    isDisabled: Boolean= false,
    isOutlined: Boolean= false

) {
    val buttonProps: ButtonProps= getButtonPropsForVariant(variant, isDisabled, isOutlined)

    Button(
        onClick = clickHandler,
        shape = RoundedCornerShape(35.dp),
        colors = buttonProps.colors,
        border = buttonProps.border,
        modifier = buttonProps.modifier,
        contentPadding = buttonProps.contentPadding
    ) {

        //* leading icon */
        if(leadingDrawableResourceID != null) {

            Image(
                painter = painterResource(id = leadingDrawableResourceID),
                contentDescription = "",
                modifier = Modifier.height(height = 18.dp)
            )

            Spacer(modifier = Modifier.width(width = 10.dp))
        }

        Text(
            text = buttonText,
            color = Color.Black,
            lineHeight = 130.sp,
            style = AppTheme.typography.boldButtonText
        )

        //* trailing icon */
        if(trailingDrawableResourceID != null) {

            Spacer(modifier = Modifier.width(width = 10.dp))

            Image(
                painter = painterResource(id = trailingDrawableResourceID),
                contentDescription = "",
                modifier = Modifier.height(height = 25.dp)
            )
        }
    }
}