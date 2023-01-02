package com.twitterclone.screens.enterRegistrationPassword

import androidx.compose.foundation.Image
import androidx.compose.foundation.layout.*
import androidx.compose.material.Divider
import androidx.compose.material.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.unit.dp
import com.twitterclone.R
import com.twitterclone.components.button.ButtonComponentView
import com.twitterclone.components.formField.FormFieldComponentView
import com.twitterclone.theme.AppTheme

@Composable
fun EnterRegistrationPasswordScreenView( ) {

    Column(modifier = Modifier.fillMaxSize( )) {

        Column(modifier = Modifier.padding(horizontal = 16.dp, vertical = 10.dp)) {

            Row(
                modifier= Modifier.fillMaxWidth( ),
                horizontalArrangement = Arrangement.Center
            ) {
                Image(
                    painter = painterResource(id = R.drawable.twitter_logo),
                    contentDescription = "",
                    modifier = Modifier.width(30.dp)
                )
            }

            Spacer(modifier = Modifier.height(100.dp))

            Text(text = "You will need a password", style = AppTheme.typography.extraBoldHeading)

            Spacer(modifier = Modifier.height(20.dp))

            Text(
                text = "The password should have at-least 8 characters",
                style = AppTheme.typography.smallSizedParagraphText.copy(color = AppTheme.colors.text)
            )

            Spacer(modifier = Modifier.height(17.5f.dp))

            FormFieldComponentView(
                labelText = "Secure password",
                leadingIconResourceId = R.drawable.password,
                isSensitive = true
            )
        }

        Spacer(modifier = Modifier.weight(1f))

        Divider(color = AppTheme.colors.lightBorder)

        Box(
            modifier = Modifier
                .fillMaxWidth( )
                .padding(vertical = 20.dp, horizontal = 16.dp),
            contentAlignment = Alignment.CenterEnd
        ) {
            ButtonComponentView(
                buttonText = "Next",
                clickHandler = { },
                trailingDrawableResourceID = R.drawable.arrow_front_circle_outline,
                isDisabled = true
            )
        }
    }
}