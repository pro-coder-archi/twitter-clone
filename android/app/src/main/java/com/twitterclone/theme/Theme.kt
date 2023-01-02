package com.twitterclone.theme

import android.content.res.Configuration
import androidx.compose.material.ProvideTextStyle
import androidx.compose.material.Scaffold
import androidx.compose.runtime.Composable
import androidx.compose.runtime.ReadOnlyComposable
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.platform.LocalConfiguration
import androidx.compose.ui.text.TextStyle

object AppTheme {

    val colors: ThemeColors
        get( )= getDarkThemeColors( )

    val typography: ThemeTypography
        get( )= ThemeTypography( )

    val screenDimensions: Configuration
        @Composable
        @ReadOnlyComposable //* denotes that this composable will rarely affect the ui. This annotation improves performance
        get( )= LocalConfiguration.current

}

// TODO: implement light theme and functionality to switch between themes

@Composable
fun AppTheme(content: @Composable ( ) -> Unit) {
    Scaffold(backgroundColor = AppTheme.colors.background) {

        content( )

    }
}