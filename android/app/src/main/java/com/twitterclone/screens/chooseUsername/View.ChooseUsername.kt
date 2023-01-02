package com.twitterclone.screens.chooseUsername

import androidx.compose.foundation.Image
import androidx.compose.foundation.layout.*
import androidx.compose.material.Divider
import androidx.compose.material.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.unit.dp
import com.twitterclone.R
import com.twitterclone.components.button.ButtonComponentView
import com.twitterclone.components.formField.FormFieldComponentView
import com.twitterclone.theme.AppTheme

@Composable
fun ChooseUsernameView( ) {

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

                Spacer(modifier = Modifier.width(30.dp))
            }

            Spacer(modifier = Modifier.height(100.dp))

            Text(text = "What should we call you ?", style = AppTheme.typography.extraBoldHeading)

            Spacer(modifier = Modifier.height(20.dp))

            Text(
                text = "Your @username is unique. You can always change it later.",
                style = AppTheme.typography.regularSizedParagraphText.copy(color = AppTheme.colors.text)
            )

            Spacer(modifier = Modifier.height(17.5f.dp))

            FormFieldComponentView(
                labelText = "Preferred username",
                leadingIconResourceId = R.drawable.user_circle_light
            )
        }

        Spacer(modifier = Modifier.weight(1f))

        Divider(color = AppTheme.colors.lightBorder)

        Row(
            modifier = Modifier
                .fillMaxWidth( )
                .padding(vertical = 20.dp, horizontal = 16.dp)
        ) {
            ButtonComponentView(
                buttonText = "Skip for now",
                clickHandler = { },
                isOutlined = true
            )
            
            Spacer(modifier = Modifier.weight(1f))

            ButtonComponentView(
                buttonText = "Next",
                clickHandler = { },
                trailingDrawableResourceID = R.drawable.arrow_front_circle_outline,
                isDisabled = true
            )
        }
    }

}