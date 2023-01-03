package com.twitterclone.screens.verifyEmailForRegistration

import androidx.compose.foundation.Image
import androidx.compose.foundation.layout.*
import androidx.compose.material.Divider
import androidx.compose.material.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.text.SpanStyle
import androidx.compose.ui.text.buildAnnotatedString
import androidx.compose.ui.text.style.TextAlign
import androidx.compose.ui.text.withStyle
import androidx.compose.ui.unit.dp
import com.twitterclone.R
import com.twitterclone.components.button.ButtonComponentView
import com.twitterclone.components.formField.FormFieldComponentView
import com.twitterclone.theme.AppTheme

@Composable
fun VerifyEmailForRegistrationView( ) {

    Column(modifier = Modifier.fillMaxSize( )) {

        Column(modifier = Modifier.padding(horizontal = 16.dp, vertical = 10.dp)) {

            Row(
                modifier= Modifier.fillMaxWidth( ),
                horizontalArrangement = Arrangement.SpaceBetween
            ) {
                Image(
                    painter = painterResource(id = R.drawable.arrow_back_circle_outline),
                    contentDescription = "",
                    modifier = Modifier.width(30.dp)
                )

                Image(
                    painter = painterResource(id = R.drawable.twitter_logo),
                    contentDescription = "",
                    modifier = Modifier.width(30.dp)
                )

                Spacer(modifier = Modifier.width(30.dp))
            }

            Spacer(modifier = Modifier.height(100.dp))

            Text(text = "We sent you a code", style = AppTheme.typography.extraBoldHeading)

            Spacer(modifier = Modifier.height(20.dp))

            Text(
                text = buildAnnotatedString
                    {
                        append("Enter the code to verify ")

                        withStyle(style = SpanStyle(color = AppTheme.colors.contrast))
                            {append("archismanmridha12345@gmail.com")}
                    },
                style = AppTheme.typography.smallSizedParagraphText.copy(color = AppTheme.colors.text)
            )

            Spacer(modifier = Modifier.height(17.5f.dp))

            Text(
                text = "The code will be valid for next 20 minutes",
                style = AppTheme.typography.smallSizedParagraphText.copy(color = AppTheme.colors.text)
            )

            Spacer(modifier = Modifier.height(17.5f.dp))

            FormFieldComponentView(
                labelText = "Verification code",
                leadingIconResourceId = R.drawable.otp
            )

            Text(
                text = "02:00",
                style = AppTheme.typography.smallSizedParagraphText,
                modifier = Modifier.fillMaxWidth( ),
                textAlign = TextAlign.End
            )

            Spacer(modifier = Modifier.height(10f.dp))

            Text(
                text = "Resend verification code",
                style = AppTheme.typography.regularSizedParagraphText.copy(color = AppTheme.colors.error.copy(alpha = 0.5f))
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