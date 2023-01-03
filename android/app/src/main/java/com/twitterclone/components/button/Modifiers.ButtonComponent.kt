package com.twitterclone.components.button

import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.alpha

fun Modifier.buttonVariant(variant: ButtonComponentVariants, isDisabled: Boolean): Modifier {

    //* if the button variant is wide then fill maximum width */
    val modifier=
        if(variant == ButtonComponentVariants.WIDE)
            fillMaxWidth( )
        else Modifier

    //* if button is disabled then reduce transparency to 50% */
    return if(isDisabled) modifier.alpha(0.5f)
        else modifier
}