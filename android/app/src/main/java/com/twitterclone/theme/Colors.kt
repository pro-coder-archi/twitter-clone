package com.twitterclone.theme

import androidx.compose.ui.graphics.Color

private val primaryColor= Color(0xFF1D9BF0)
private val backgroundColor= Color(0XFF000000)
private val errorColor= Color(0xFFFF453A)
private val textColor= Color(250, 250, 250, 75)
private val lightBorderColor= Color(250, 250, 250, 20)

class ThemeColors(

    var primary: Color,
    var background: Color,
    var error: Color,
    var text: Color,
    var lightBorder: Color,

    var isForDarkTheme: Boolean

)

fun getDarkThemeColors(

    primary: Color= primaryColor,
    background: Color= backgroundColor,
    error: Color= errorColor,
    text: Color= textColor,
    lightBorder: Color= lightBorderColor

) =
    ThemeColors(
        primary= primary,
        background= background,
        error= error,
        text = text,
        lightBorder= lightBorder,

        isForDarkTheme= true
    )