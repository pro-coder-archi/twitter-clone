package com.twitterclone.components.button

import androidx.compose.foundation.layout.PaddingValues
import androidx.compose.material.ButtonColors
import androidx.compose.material.ButtonDefaults
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.unit.dp
import com.twitterclone.theme.AppTheme

data class ButtonProps(

    val colors: ButtonColors,
    val modifier: Modifier,
    val contentPadding: PaddingValues
)

@Composable
fun getButtonPropsForVariant(variant: ButtonComponentVariants, isDisabled: Boolean, isOutlined: Boolean): ButtonProps {

    //* decide the button background and content color based on the variant is outlined or not */
    val buttonColor= ButtonDefaults.buttonColors(

        backgroundColor =
            if(isOutlined) Color.Transparent
                else AppTheme.colors.contrast,

        contentColor =
            if(isOutlined) AppTheme.colors.contrast
                else AppTheme.colors.background
    )

    return when(variant) {

        ButtonComponentVariants.WIDE ->
            ButtonProps(
                colors = buttonColor,
                modifier = Modifier.buttonVariant(ButtonComponentVariants.WIDE, isDisabled),
                contentPadding = PaddingValues(vertical = 12.5f.dp, horizontal = 15.dp)
            )

        ButtonComponentVariants.CHIP ->
            ButtonProps(
                colors = buttonColor,
                modifier = Modifier.buttonVariant(ButtonComponentVariants.CHIP, isDisabled),
                contentPadding = PaddingValues(vertical = 5.dp, horizontal = 15.dp)
            )

        else ->
            ButtonProps(
                colors = buttonColor,
                modifier = Modifier.buttonVariant(ButtonComponentVariants.DEFAULT, isDisabled),
                contentPadding = PaddingValues(vertical = 7.5f.dp, horizontal = 15.dp)
            )
    }
}