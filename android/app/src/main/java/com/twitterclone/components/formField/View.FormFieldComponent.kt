package com.twitterclone.components.formField

import androidx.compose.foundation.Image
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.*
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.foundation.text.KeyboardOptions
import androidx.compose.material.OutlinedTextField
import androidx.compose.material.Text
import androidx.compose.material.TextFieldDefaults
import androidx.compose.runtime.*
import androidx.compose.ui.Modifier
import androidx.compose.ui.focus.onFocusChanged
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.text.input.KeyboardType
import androidx.compose.ui.text.input.PasswordVisualTransformation
import androidx.compose.ui.text.input.VisualTransformation
import androidx.compose.ui.text.style.TextAlign
import androidx.compose.ui.unit.dp
import com.twitterclone.R
import com.twitterclone.theme.AppTheme

fun getBackgroundColor(isError: Boolean): Color {
    return if(isError)
        AppTheme.colors.error.copy(alpha = 0.15f)
    else Color.Transparent
}

@Composable
fun FormFieldComponentView(

    labelText: String,
    leadingIconResourceId: Int,
    isSensitive: Boolean= false

) {
    var currentValue: String by remember{ mutableStateOf("") }

    var isFocused: Boolean by remember{ mutableStateOf(false) }
    var isTextVisible: Boolean by remember{ mutableStateOf(!isSensitive) }

    OutlinedTextField(
        value = currentValue,
        onValueChange = { currentValue= it },
        keyboardOptions = KeyboardOptions(keyboardType = KeyboardType.Email),

        leadingIcon = {
            Image(
                painter = painterResource(id = leadingIconResourceId),
                contentDescription = "",
                modifier = Modifier.height(25.dp)
            )
        },

        trailingIcon = {
            if(currentValue != "" && isFocused)
                Row {

                    Image(
                        painter = painterResource(id = R.drawable.close),
                        contentDescription = "",
                        modifier = Modifier
                            .height(20.dp)
                            .clickable { currentValue= "" },
                    )

                    // for sensitive form field
                    if(isSensitive) {
                        Spacer(modifier = Modifier.width(10.dp))

                        Image(
                            painter = painterResource(id = if(isTextVisible) R.drawable.striked_eye else R.drawable.eye),
                            contentDescription = "",
                            modifier = Modifier
                                .height(20.dp)
                                .clickable { isTextVisible= !isTextVisible },
                        )

                        Spacer(modifier = Modifier.width(10.dp))
                    }
                }
        },

        label = {
            if(isFocused)
                Text(
                    text = labelText,
                    style = AppTheme.typography.smallSizedPlaceholderText.copy(color = AppTheme.colors.text)
                )

            else Text(
                text = labelText,
                style = AppTheme.typography.smallSizedSemiBoldText.copy(color = AppTheme.colors.text)
            )
        },

        modifier = Modifier
            .fillMaxWidth( )
            .onFocusChanged { isFocused= it.isFocused }
            .height(60.dp),

        visualTransformation = if(isTextVisible) VisualTransformation.None else PasswordVisualTransformation( ),

        shape = RoundedCornerShape(15.dp),
        textStyle = AppTheme.typography.smallSizedSemiBoldText,

        colors = TextFieldDefaults.outlinedTextFieldColors(
            unfocusedBorderColor = AppTheme.colors.text,
            focusedBorderColor = AppTheme.colors.primary,
            errorBorderColor = AppTheme.colors.error,
            cursorColor = AppTheme.colors.primary,
            textColor = AppTheme.colors.contrast,
//            backgroundColor = getBackgroundColor(true)
        ),

//        isError = true
    )

//    Spacer(modifier = Modifier.height(5f.dp))
//
//    Text(
//        text = "Full name cannot be empty",
//        style = AppTheme.typography.smallSizedParagraphText.copy(color = AppTheme.colors.error),
//        modifier = Modifier.fillMaxWidth( ),
//        textAlign = TextAlign.End
//    )

    Spacer(modifier = Modifier.height(12.5.dp))

}